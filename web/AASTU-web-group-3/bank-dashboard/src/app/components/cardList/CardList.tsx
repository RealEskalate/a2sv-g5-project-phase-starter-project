import React from "react";
import { useState } from "react";
// import { CardData } from "@/types/cardData";
import { Card } from "@/lib/redux/types/cards";
import CardBox from "@/app/components/cardBox";
import {
  useDeleteCardMutation,
  useGetCardsQuery,
} from "@/lib/redux/api/cardsApi";
import Loading from "@/app/loading";

function formatCardNumber(cardNumber: string): string {
  const start = cardNumber.slice(0, 4);
  const end = cardNumber.slice(-4);
  return `${start} **** **** ${end}`;
}

const colorOptions = [
  { svgColor: "#396AFF", svgBgColor: "#E7EDFF" },
  { svgColor: "#FF82AC", svgBgColor: "#FFE0EB" },
  { svgColor: "#FFD700", svgBgColor: "#FFF5D9" },
];

const CardListPage = () => {
  const [page, setPage] = useState(0);
  const { data, error, isLoading } = useGetCardsQuery({ page, size: 3 });

  const totalPages = data?.totalPages || 1;

  const handlePreviousPage = () => {
    if (page > 0) {
      setPage(page - 1);
    }
  };

  const handleNextPage = () => {
    if (page < totalPages - 1) {
      setPage(page + 1);
    }
  };

  if (isLoading) return <Loading />;
  if (error) return <div>Error fetching cards</div>;

  return (
    <div className="card-list-container flex flex-col justify-between h-full bg-white dark:bg-darkComponent">
    <div className="card-list flex-grow">
      {data?.content.map((card: Card, index: number) => {
        const colorOption = colorOptions[index % colorOptions.length]; // Cycle through color options
  
        return (
          <CardBox
            key={index}
            cardType={card.cardType}
            bank="CBE Bank"
            detailsLink="/creditcardpage/xyz"
            svgColor={colorOption.svgColor}
            svgBgColor={colorOption.svgBgColor}
            cardNumber={formatCardNumber(card.semiCardNumber)}
            NamainCard={card.cardHolder}
          />
        );
      })}
    </div>
  

      {/* <div className="pagination-controls flex justify-center items-center mt-4">
        <button
          onClick={handlePreviousPage}
          disabled={page === 0}
          className="btn"
        >
          Previous
        </button>
        <span className="mx-2">
          Page {page + 1} of {totalPages}
        </span>
        <button
          onClick={handleNextPage}
          disabled={page >= totalPages - 1}
          className="btn"
        >
          Next
        </button>
      </div> */}
    </div>
  );
};
export default CardListPage;
