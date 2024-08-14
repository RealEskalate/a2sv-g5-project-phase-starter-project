import { Inter } from "next/font/google";

const inter = Inter({
  subsets: ["latin"],
});

interface PersonCardType {
  imageLink: string;
  fullName: string;
  jobTitle: string;
}

const PersonCard = ({ imageLink, fullName, jobTitle }: PersonCardType) => {
  return (
    <div className="flex flex-col gap-3 items-center">
      <div>
        <img
          src={imageLink}
          alt="person image"
          className="w-[70px] h-[70px] rounded-full"
        />
      </div>
      <div className="flex flex-col items-center">
        <div className={`text-[#232323] font-normal ${inter.className}`}>
          {fullName}
        </div>
        <div className={`text-[#718EBF] font-normal ${inter.className}`}>
          {jobTitle}{" "}
        </div>
      </div>
    </div>
  );
};

export default PersonCard;
