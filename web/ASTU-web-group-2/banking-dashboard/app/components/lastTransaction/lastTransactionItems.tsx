type Item = {
  image: string;
  title: string;
  date: string;
  category: string;
  pass: string;
  status: string;
  amount: string;
};

export const items: Item[] = [
  {
    image: "/assets/lastTransaction/spot-sub.svg",
    title: "Spotify Subscription",
    date: "25 Jan 2021",
    category: "Shopping",
    pass: "1234 ****",
    status: "Pending",
    amount: "-$150",
  },

  {
    image: "/assets/lastTransaction/settings.svg",
    title: "Mobile Service",
    date: "25 Jan 2021",
    category: "Service",
    pass: "1234 ****",
    status: "Completed",
    amount: "-$340",
  },

  {
    image: "/assets/lastTransaction/user.svg",
    title: "Emily Watson",
    date: "25 Jan 2021",
    category: "Transfer",
    pass: "1234 ****",
    status: "Completed",
    amount: "+$780",
  },
];
