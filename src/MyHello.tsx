type MyHelloProps = {
  myName: string
}

export default function MyHello({ myName }: MyHelloProps) {
  return (
    <div>こんにちは{myName}</div>
  )
}
