import AddNewCardForm from "../Form/AddNewCardForm";

const AddNewCard = () => {
  return (
    <div className=" w-full md:w-8/12">
      <h1 className="text-[#333B69] py-2 font-semibold">Add New Card</h1>
      <div className="bg-white p-5 rounded-2xl">
        <p className="text-15px py-2 text-blue-steel">
          Credit Card generally means a plastic card issued by Scheduled
          Commercial Banks assigned to a Cardholder, with a credit limit, that
          can be used to purchase goods and services on credit or obtain cash
          advances.
        </p>
        <AddNewCardForm />
      </div>
    </div>
  );
};

export default AddNewCard;
