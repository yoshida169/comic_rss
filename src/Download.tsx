import dl_icon from '../public/download.png'

export default function Download({ isbn }: { isbn: string }) {
  return (
    <a href={`https://wings.web-deli.com/books/${isbn}.jpg`}>
      <img src={dl_icon} alt="sample download"/>
    </a>
  );
}
