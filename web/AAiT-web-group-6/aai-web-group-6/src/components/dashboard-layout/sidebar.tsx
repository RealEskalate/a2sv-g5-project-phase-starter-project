'use client'
import React from 'react'
import SideBarButton from '@/components/dashboard-layout/sidebar_button'
import { sidebarLinks } from '@/constants/Layout/lay_constant'
import { usePathname } from 'next/navigation'

type Props = {}

export default function SideBar({ }: Props) {
  const currentPath = usePathname();
  return (
    <div className='flex flex-col gap-y-1'>
      {sidebarLinks.map((value, ind) => {
        return <SideBarButton key={ind} Icon={value.icon} isActive={currentPath == value.route} route={value.route} title={value.title} />
      })}
    </div>
  )
}