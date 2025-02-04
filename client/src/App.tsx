import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
// import './App.css'
import { Stack, Container } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'

export const BASE_URL = "http://localhost:5000/api"
function App() {
  const [count, setCount] = useState(0)

  return (
    <Stack h="100vh">
    <Navbar />
    <Container>
      <TodoForm />
      <TodoList />
    </Container>
    </Stack>
  )
}

export default App
