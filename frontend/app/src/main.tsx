import React from 'react'
import ReactDOM from 'react-dom/client'
import { ChakraProvider, Container, ColorModeScript } from '@chakra-ui/react'
import theme from './theme'
import App from './App.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ChakraProvider theme={theme}>
      <ColorModeScript initialColorMode={theme.config.initialColorMode} />
      <Container maxW='4xl' centerContent>
        <App />
      </Container>
    </ChakraProvider>
  </React.StrictMode>,
)
