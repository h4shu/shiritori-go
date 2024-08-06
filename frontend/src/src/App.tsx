import { useState } from 'react'
import useSWR from 'swr'
import { VStack, Input, Text, Card, Alert, AlertIcon } from '@chakra-ui/react'
import LastWord from './components/shiritori/LastWord'

const backendBaseURL = import.meta.env.VITE_BACKEND_BASEURL;

const postFetcher = (url: string, data: string) => fetch(
  url, {
      mode: "cors",
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: data
  })

const getFetcher = (url: string) => fetch(
  url, {
      mode: "cors",
      method: "GET",
      headers: {
          "Content-Type": "application/json",
      }
  }
).then(res => res.json())

function App() {
  const [errMsg, setErrMsg] = useState("")

  const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault()
    const form = new FormData(event.currentTarget)
    const word = form.get("word") as string
    const data = JSON.stringify({
      word: word
    })
    postFetcher(`${backendBaseURL}/wc`, data).then(res => {
      if (!res.ok) {
        res.json().then((json) => {
            if (json.message)
              setErrMsg(json.message)
            else
              setErrMsg(res.status + ' ' + res.statusText)
        })
      }
      else
        location.reload()
    }).catch(error => {
        setErrMsg(error.toString())
    })
  }

  let alertMsg
  if (errMsg)
    alertMsg = <Alert status='error'><AlertIcon />{errMsg}</Alert>

  const { data, error, isLoading } = useSWR(`${backendBaseURL}/wc`, getFetcher)
  let history
  let historyLen = "0"
  if (error)
    history = <Text>failed to load</Text>
  else if (isLoading)
    history = <Text>loading...</Text>
  if (data && data.wordchain) {
    history = data.wordchain.map((word: string) => {
      return (
        <Card>{word}</Card>
      )
    })
    historyLen = data.len
  }

  return (
    <VStack>
      <Text>ことばのかず：{historyLen}</Text>
      <LastWord />
      <form onSubmit={handleSubmit}>
        <Input name='word' placeholder='つぎのことば'/>
      </form>
      {alertMsg}
      {history}
    </VStack>
  )
}

export default App
