import Head from "next/head";

export default function Home() {
	return (
		<>
			<Head>
				<title>Hello Next</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<div className="grid justify-center content-center h-screen">
				<p>Hello, World!</p>
			</div>
		</>
	);
}
