import './App.css'
import MyHello from './MyHello'

function App() {
  return (
    <div className="App">
      <MyHello myName="太郎" />
      <button type="submit" onClick={() => getComics()}>
        DB通信
      </button>
    </div>
  )
}

export default App

function getComics() {
  console.log('get')
}