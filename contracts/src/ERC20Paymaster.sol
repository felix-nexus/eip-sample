// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/utils/SafeERC20.sol";

import "@account-abstraction-0.8.0/contracts/core/BasePaymaster.sol";

contract ERC20Paymaster is BasePaymaster {
    using SafeERC20 for IERC20;

    IERC20 public immutable token;
    uint256 public tokenPerGas;
    mapping(bytes32 => bool) public usedUserOps;

    constructor(IEntryPoint _entryPoint, IERC20 _token, uint256 _tokenPerGas) BasePaymaster(_entryPoint) {
        token = _token;
        tokenPerGas = _tokenPerGas;
    }

    function _validatePaymasterUserOp(PackedUserOperation calldata userOp, bytes32 userOpHash, uint256 maxCost)
        internal
        override
        returns (bytes memory context, uint256 validationData)
    {
        if (usedUserOps[userOpHash]) {
            revert("UserOp already used");
        }
        usedUserOps[userOpHash] = true;

        address sender = userOp.sender;
        uint256 requiredTokenAmount = (maxCost * tokenPerGas) / 1e18;

        require(token.balanceOf(sender) >= requiredTokenAmount, "Insufficient token");

        return (abi.encode(sender, requiredTokenAmount), 0);
    }

    function _postOp(PostOpMode mode, bytes calldata context, uint256 actualGasCost, uint256 actualUserOpFeePerGas)
        internal
        override
    {
        (address sender, uint256 preCharged) = abi.decode(context, (address, uint256));

        uint256 gasCostInETH = actualGasCost * actualUserOpFeePerGas;
        uint256 tokenAmount = (gasCostInETH * tokenPerGas) / 1e18;

        if (mode == PostOpMode.opReverted) {
            // 특별한 정책 적용 가능 (e.g., 반만 징수)
            tokenAmount = tokenAmount / 2;
        }

        if (tokenAmount > preCharged) {
            tokenAmount = preCharged; // 과금 상한
        }

        token.safeTransferFrom(sender, address(this), tokenAmount);
    }

    function setTokenPerGas(uint256 _price) external onlyOwner {
        tokenPerGas = _price;
    }

    function withdrawToken(address to) external onlyOwner {
        uint256 bal = token.balanceOf(address(this));
        require(token.transfer(to, bal), "Withdraw failed");
    }
}
