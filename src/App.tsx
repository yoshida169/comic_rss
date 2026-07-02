import './App.css'
import MyHello from './MyHello'
import StateBasic from './StateBasic'
import books from './Books'
import ForList from './ForList'

function App() {
  return (
    <>
      <ForList src={books} />
      <div className="App">
        <MyHello myName="太郎" />
        <button type="submit" onClick={() => getComics()}>
          DB通信
        </button>
        <StateBasic init={0} />
      </div>
    </>
  )
}

export default App

function getComics() {
  console.log('get')
}