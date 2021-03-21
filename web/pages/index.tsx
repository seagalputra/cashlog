import { NextPage } from "next";
import Head from "next/head";
import Banner from "../components/Banner";

const ListItem = () => {
	return (
		<div className="flex justify-between items-center max-w-xl mx-auto my-12">
			<div className="flex items-center">
				<div
					className="flex justify-center items-center rounded-full border-2"
					style={{
						borderColor: "#18191F",
						height: "72px",
						width: "72px",
						backgroundColor: "#FFD465",
					}}
				>
					<p
						className="font-extrabold text-2xl"
						style={{
							fontFamily: "Montserrat, sans-serif",
							color: "#18191F",
						}}
					>
						F
					</p>
				</div>
				<div className="ml-6">
					<p
						className="font-extrabold text-2xl"
						style={{
							fontFamily: "Montserrat, sans-serif",
							color: "#18191F",
						}}
					>
						Food
					</p>
					<p
						className="font-medium"
						style={{
							fontFamily: "Montserrat, sans-serif",
							color: "#18191F",
						}}
					>
						Pizza, Burger, Chinese...
					</p>
				</div>
			</div>
			<p
				className="font-bold text-lg"
				style={{
					fontFamily: "Montserrat, sans-serif",
					color: "#9FA4B4",
				}}
			>
				time
			</p>
		</div>
	);
};

const Home: NextPage = () => {
	return (
		<>
			<Head>
				<title>Home | Cashlog</title>
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<main>
				<Banner />
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
				{Array(4)
					.fill(<ListItem />)
					.map((item) => item)}
			</main>
		</>
	);
};

export default Home;
