/**
 * generame el tsx para un boton. Quiero poder parametrizar el nombre, el largo pero con opciones s, m, l, y el tipo de boton para color si es cancelar rojo, si es guardar verde, si es continuar azul
 */
import React from 'react';
import Button from '@mui/material/Button';
import {styled} from '@mui/material/styles';

interface CustomButtonProps {
    id:string;
    label: string;
    size?: 's' | 'm' | 'l'| 'xl'| 'xxl';
    type?: 'cancel' | 'save' | 'continue';
    onClick?: () => void;
}

const StyledButton = styled(Button)<{ btnType: string, btnSize: string }>(({theme, btnType, btnSize}) => ({
    backgroundColor:
        btnType === 'cancel' ? theme.palette.error.main :
            btnType === 'save' ? theme.palette.success.main :
                btnType === 'continue' ? theme.palette.primary.main : theme.palette.grey[500],
    color: theme.palette.common.white,
    margin: 5,
    minWidth:   getSize(btnSize),
    '&:hover': {
        backgroundColor:
            btnType === 'cancel' ? theme.palette.error.dark :
                btnType === 'save' ? theme.palette.success.dark :
                    btnType === 'continue' ? theme.palette.primary.dark : theme.palette.grey[700],
    },
}));

const ButtonAtom: React.FC<CustomButtonProps> = ({id,label, size = 'm', type = 'continue', onClick}) => {
    return (
        <StyledButton id={id} btnType={type} btnSize={size} variant="contained" onClick={onClick}>
            {label}
        </StyledButton>
    );
};

export default ButtonAtom;
