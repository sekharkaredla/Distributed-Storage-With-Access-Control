pragma solidity >=0.4.22 <0.6.0;
contract StoringJWT {
 address owner;
 event StoringJWTEvent(address _sender, bytes32 _JWT);

 constructor(bytes32 _JWT) public {
     owner = msg.sender;
     emit StoringJWTEvent(msg.sender, _JWT);
 }

 function getOwner() public returns (address) {
   return owner;
 }

}