import useSWR from 'swr'
import { Text } from '@chakra-ui/react'

const backendBaseURL = import.meta.env.VITE_BACKEND_BASEURL;

const getFetcher = (url: string) => fetch(
  url, {
      mode: "cors",
      method: "GET",
      headers: {
          "Content-Type": "application/json",
      }
  }
).then(res => res.json())

function LastWord() {
  const { data, error, isLoading } = useSWR(`${backendBaseURL}/wc/last`, getFetcher)
  let word
  if (error)
    word = <Text>failed to load</Text>
  else if (isLoading)
    word = <Text>loading...</Text>
  else if (data)
    word = <Text>まえのことば: {data.word} </Text>

  return (
    <>{word}</>
  )
}

export default LastWord
