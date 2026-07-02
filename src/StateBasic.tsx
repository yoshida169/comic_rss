import { useState } from 'react';

export default function StateBasic({ init }: { init: number }) {
  const [count, setCount] = useState(init)

  const handleClick = () => {
    setCount(count + 1)
  }

  return (
    <>
      <button onClick={handleClick}>+1</button>
      <p>Count: {count}</p>
    </>
  )
}