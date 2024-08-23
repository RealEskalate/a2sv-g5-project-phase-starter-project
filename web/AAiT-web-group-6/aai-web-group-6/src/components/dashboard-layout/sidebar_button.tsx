import Link from 'next/link';
import React from 'react'
import { SvgIconProps } from '@mui/material/SvgIcon/SvgIcon';

type Props = {
    route: string;
    Icon: (props: SvgIconProps) => React.JSX.Element;
    title: string;
    isActive: boolean;
}
const ActiveClass = "text-[#1814F3]";
const InActiveClass = "text-[#B1B1B1]";
export default function SideBarButton({ Icon, route, title, isActive }: Props) {
    return (
        <Link href={route} className='flex  items-center'>
            <div className={`w-[6px] h-[60px] rounded-r-xl ${isActive ? 'bg-[#1814F3]' : "bg-white"}`}></div>
            <div className={`flex gap-x-5 pl-6  py-2 ${isActive ? ActiveClass : InActiveClass}`} >
                <Icon className={isActive ? ActiveClass : InActiveClass} />
                <p className='font-medium text-[18px]'>{title}</p>
            </div>
        </Link>
    )
}