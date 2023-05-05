pragma solidity ^0.5.16;

contract HelloWorld {
    string public message;

    constructor(string memory _message) public {
        message = _message;
    }

    function getMessage() public view returns (string memory) {
        return message;
    }

    function setMessage(string memory _message) public {
        message = _message;
    }
}