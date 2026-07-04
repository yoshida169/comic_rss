import type { Book } from './Books'

export default function ForItem({ book }: { book: Book }) {
  return (
    <>
      <dt>
        <a href={`https://wings.web-deli.com/books/${book.isbn}.jpg`}>
          {book.title} ({book.price}円)
        </a>
      </dt>
      <dd>{book.summary}</dd>
    </>
  )
}
