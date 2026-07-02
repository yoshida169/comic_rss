export type Book = {
  isbn: string;
  title: string;
  price: number;
  summary: string;
};

const books: Book[] = [
  {
    isbn: 'aaa',
    title: '本のタイトル',
    price: 1000,
    summary: '本のサマリ',
  },
  {
    isbn: 'bbb',
    title: '本のタイトル2',
    price: 2000,
    summary: '本のサマリ2',
  },
  {
    isbn: 'bbb',
    title: '本のタイトル3',
    price: 3000,
    summary: '本のサマリ3',
  },
]
export default books;