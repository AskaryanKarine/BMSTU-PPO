import React from 'react';
import './styles.css'
import {Link} from "react-router-dom";

interface HeaderProps {
    children: React.ReactNode
}

export function Header({children}:HeaderProps) {
    return (
        <div className="header">
            <Link to="/" className='w-[20%]'>
                <h1 className='text-3xl font-serif font-bold text-center'>DoramaSet</h1>
            </Link>
            {children}
        </div>
    )
}