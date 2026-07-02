import React from "react";

type Book = {
  isbn: string;
  title: string;
  price: number;
  summary: string;
};

export default function ForList({ src }: { src: Book[] }) {
  return (
      <dl>
        {src.map((elem) => (
          <React.Fragment key={elem.isbn}>
            <dt>
              <a href={`https://wings.web-deli.com/books/${elem.isbn}.jpg`}>
                {elem.title} ({elem.price}円)
              </a>
            </dt>
            <dd>{elem.summary}</dd>
          </React.Fragment>
        ))}
      </dl>
  );
}