import React from "react";

export default function ForList( {src} ){
  return(
    <dl>
      {
      src.map(elem => (
        <>
          <dt>
            <a href={`https://wings.web-deli.com/books/${elem.isbn}.jpg`}
            {elem.title} ({elem.price}円)
          <dt/>
          <dd>{elem.summary}</dd>
        </>
      ))}
    <dl/>
  );
}