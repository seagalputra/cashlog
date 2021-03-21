import BalanceCard from "./BalanceCard";

const Banner = () => {
	return (
		<>
			<div
				className="pb-4 pt-16 flex flex-col justify-center items-center"
				style={{
					backgroundColor: "#00C6AE",
				}}
			>
				<img src="/person.svg" alt="Two person" />
				<BalanceCard />
			</div>
		</>
	);
};

export default Banner;
