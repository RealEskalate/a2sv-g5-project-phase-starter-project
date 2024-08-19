import React from 'react';
import { colors } from '@/constants/index';


type StockData = {
    slNo: string;
    name: string;
    price: string;
    return: string;
};

// const data: StockData[] = [
//     { slNo: '04.', name: 'Nokia', price: '$940', return: '+2%', returnColor: 'text-green-500' },
// ];

interface props{
    items:StockData[]
}
const TrendingStock = ({items}:props) => {
   

    return (
        <div className="w-[100%] p-4 h-[343px] rounded-3xl bg-white overflow-y-auto ">
            <table className="w-full border-collapse">
                <thead >
                    <tr>
                        <th className={`text-left py-2 px-3 border-b border-gray-300 ${colors.textgray}`}>SL No</th>
                        <th className={`text-left py-2 px-3 border-b border-gray-300 ${colors.textgray}`}>Name</th>
                        <th className={`text-left py-2 px-3 border-b border-gray-300 ${colors.textgray}`}>Price</th>
                        <th className={`text-left py-2 px-3 border-b border-gray-300 ${colors.textgray}`}>Return</th>
                    </tr>
                </thead>
                <tbody>
                    {items.map((item, index) =>
                        <tr key={item.slNo}>
                            <td className="py-2 px-3  text-gray-800">{item.slNo}</td>
                            <td className="py-2 px-3  text-gray-800">{item.name}</td>
                            <td className="py-2 px-3  text-gray-800">{item.price}</td>
                            {
                                item.return.includes('+') ? <td className={`py-2 px-3 text-green-500`}>{item.return}</td> :
                                    <td className={`py-2 px-3 text-red-500`}>{item.return}</td>
                            }
                        </tr>
                    )} 
                </tbody>
            </table>
        </div>
    );
};

export default TrendingStock;
