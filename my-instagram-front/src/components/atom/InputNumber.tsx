/**
 * creame InputNumber, que pueda pasar, tamaño s, m,l , que haga validaciones de valor minimo y valor maximo sean valores opcionales que pase. label dentro del input .
 */
import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import { styled } from '@mui/material/styles';

interface InputNumberProps {
    id: string;
    label: string;
    value: number;
    onChange: (value: number) => void;
    size?: 's' | 'm' | 'l'| 'xl'| 'xxl';
    min?: number;
    max?: number;
    validate?: boolean;
}

const StyledTextField = styled(TextField)<{ inputSize: string }>(({ inputSize }) => {
    let inputWidth: string;

    inputWidth = getSize(inputSize);

    return {
        // Establecemos los márgenes de 5px alrededor del componente
        margin: 5,
        '& .MuiInputBase-root': {
            minWidth: inputWidth,
            maxWidth: inputWidth,
        },
    };
});

const InputNumber: React.FC<InputNumberProps> = ({id, label, value, onChange, size = 'm', min, max,validate }) => {
    const [error, setError] = useState<string | null>(null);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const newValue = Number(event.target.value);

        onChange(newValue);
        setError(null);
        if(validate==false){
            return
        }
        // Validación de rango
        if ((min && newValue < min) || (max && newValue > max)) {
            setError(`El valor debe estar entre ${min ?? '-' } y ${max ?? '-'}`);
        }
    };

    return (
        <StyledTextField
            id={id}
            label={label}
            variant="outlined"
            value={value}
            onChange={handleChange}
            inputSize={size}
            type="number"
            error={!!error}
            helperText={error || ' '}
            fullWidth
        />
    );
};

export default InputNumber;
