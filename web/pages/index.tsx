import Head from "next/head";
import Banner from "../components/Banner";

const Home: React.FC = () => {
	return (
		<>
			<Head>
				<title>Home | Cashlog</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<main>
				<Banner />
			</main>
			<div className="flex justify-between max-w-xl mx-auto mt-6">
				<p
					className="text-2xl"
					style={{
						fontFamily: "Montserrat, sans-serif",
						fontWeight: "bold",
						color: "#9FA4B4",
					}}
				>
					Recent Transaction
				</p>
				<button
					className="text-2xl"
					style={{
						fontFamily: "Montserrat, sans-serif",
						color: "#F95A2C",
						fontWeight: "bold",
					}}
				>
					see all
				</button>
			</div>
		</>
	);
};

export default Home;
