import {
  Person, 
  // PhoneType
} from './protobuf/person';

// @ref https://github.com/stephenh/ts-proto?tab=readme-ov-file#example-types

// const p = Person.create({
//     id:1,
//     name:"test_name",
//     phones:[{
//         number:"09...",
//         type:PhoneType.PHONE_TYPE_HOME
//     }],
//     email:"test@test.com",
//     lastUpdated:new Date(),
// });

// const encoded = Person.encode(p).finish();
// console.log("encoded:\n",encoded);

// const decoded = Person.decode(encoded);
// console.log("decoded:\n",decoded);


(async () => {
    try {
      const rawRes = await fetch('http://localhost:3000/proto');
      console.log('Status Code:', rawRes.status);

      const unreadableBytes = await rawRes.arrayBuffer();
      const readableBytes = new Uint8Array(unreadableBytes);
  
      const res =Person.decode(readableBytes);
  
        console.log(`Response is :\n`,res);
    } catch (err) {
      console.log(err.message);
    }
  })();
