import { useState } from 'react'

export default function Message() {
  const [input, setInput] = useState('')
  const [messages, setMessages] = useState<string[]>([])

  const handleSendMessage = () => {
    if (!input.trim()) return
    setMessages((prev) => [...prev, input])
    setInput('')
  }

  return (
    <div>
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
      />
      <button onClick={handleSendMessage}>送信</button>
      {messages.map((msg, i) => (
        <p key={i}>{msg}</p>
      ))}
    </div>
  )
}
