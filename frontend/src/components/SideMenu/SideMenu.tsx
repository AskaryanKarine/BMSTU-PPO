import React from 'react';
import styles from './SideMenu.module.css'
import {useAppSelector} from "../../hooks/redux";

export function SideMenu() {
    const {token} = useAppSelector(state => state.userReducer)
    return (
        <div className={styles.menu}>
            <nav className={`${styles.navigation} ${styles.menu}`}>
                <div className={styles.firstDiv}>Публичные списки</div>
                <div>Дорамы</div>
                <div>Стафф</div>
                {token.length > 0 &&
                    <>
                        <div>Мои списки</div>
                        <div>Избранное</div>
                    </>
                }
            </nav>
        </div>
    )
}