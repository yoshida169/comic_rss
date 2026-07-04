import type { Book } from './Books'
import ForItem from './ForItem'

export default function ForNest({ src }: { src: Book[] }) {
  return (
    <dl>
      {src.map(elem => (
        <ForItem key={elem.isbn} book={elem} />
      ))}
    </dl>
  )
}