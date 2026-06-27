import './App.css'

function App() {
  return (
    <div className="App">
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