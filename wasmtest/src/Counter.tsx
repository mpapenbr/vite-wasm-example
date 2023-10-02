import { useState } from "react";
import wasmMethods from "./wasmdemo/wasmloader";
function Counter() {
  const [count, setCount] = useState(0);

  return (
    <>
      <button
        onClick={() => {
          const x = wasmMethods.demoAdd(count, 2);
          // const x = globalThis.myAdd(count, 2);
          console.log(x);
          setCount(x);
          const messages = wasmMethods.getMessages();
          console.log(messages);
        }}
      >
        count is {count}
      </button>
    </>
  );
}

export default Counter;
