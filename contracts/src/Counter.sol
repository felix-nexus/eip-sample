// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {MockERC20} from "solmate-6.8.0/src/test/utils/mocks/MockERC20.sol";

contract Counter {
    event Caller(address indexed caller);

    uint256 public number;

    function setNumber(uint256 newNumber) public {
        number = newNumber;
        emit Caller(msg.sender);
    }

    function increment() public {
        number++;
        emit Caller(msg.sender);
    }
}
