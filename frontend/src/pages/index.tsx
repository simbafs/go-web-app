import Image from 'next/image'
import useSWR from 'swr'

export default function Home() {
	const { data, error } = useSWR('/api/hello', url => fetch(url).then(res => res.json()))

	return (
		<>
			<h1 className="text-4xl">Hello world!</h1>
			<Image src="/ice.jpg" alt="ice" width={500} height={500} />
			{error ? (
				<pre>Error: {JSON.stringify(error, null, 2)}</pre>
			) : data ? (
				<pre>Data: {JSON.stringify(data, null, 2)}</pre>
			) : (
				<pre>Loading...</pre>
			)}
		</>
	)
}
