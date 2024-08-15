import Image from "next/image";
import CreditCardPage from "./creditcardpage/page";
import CardBox from "./components/cardBox/page";
import DoughnutChart from "./components/charts/Doughnut/page";
import CardListPage from "./components/cardList/CardList";

export default function Home() {
  return (
    <div>
      <CreditCardPage />
      {/* <CardBox /> */}
      {/* <DoughnutChart /> */}
      {/* <CardListPage /> */}
    </div>
  );
}
