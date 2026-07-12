import type { Book } from './Books'
import Download from './Download';

export default function ForItem({ book }: { book: Book }) {
  return (
  <>
    <dt>
    <a href={`https://wings.web-deli.com/books/${book.isbn}.jpg`}>
      {book.title} ({book.price}円)
    </a>
    </dt>
    {(() => {
      if (book.download) {
        return <dd>{book.summary}<Download isbn={book.isbn} /></dd>
      } else {
        return <dd>{book.summary}</dd>
      }
    })()}
  </>
  )
}
