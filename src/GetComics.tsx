import { useEffect, useState } from 'react';

export default function GetComics() {
  const [comics, setComics] = useState([])

  useEffect(() => {
    const fetchComics = async () => {
      const response = await fetch('http://localhost:8080/')
      const data = await response.json()
      setComics(data)
    }

    fetchComics()
  }, [])

  return (
    <div>
      <button type="submit" onClick={() => {}}>漫画取得</button>
      <pre>{JSON.stringify(comics, null, 2)}</pre>
    </div>
  )
}