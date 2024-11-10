import './App.css'
import { BrowserRouter as Router } from 'react-router-dom'
import ConfigRoutes from './routes'

function App() {
  return (
    <Router future={{ v7_startTransition: true }}>
      <ConfigRoutes />
    </Router>
  )
}

export default App
