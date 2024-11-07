import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import { styled } from '@mui/material/styles';

interface InputTextProps {
    label: string;
    value: string;
    onChange: (value: string) => void;
    size?: 's' | 'm' | 'l'| 'xl'| 'xxl';
    minLength?: number;
    maxLength?: number;
    validate?: boolean;
}

const StyledTextField = styled(TextField)<{ inputSize: string }>(({ inputSize }) => {
    let inputWidth: string;

    inputWidth = getSize(inputSize);
    return {
        // Establecemos los márgenes de 5px alrededor del componente
        marginTop: 5,
        marginRight: 5,
        marginLeft: 5,
        '& .MuiInputBase-root': {
            minWidth: inputWidth,
            maxWidth: inputWidth,
        },
    };
});

const InputText: React.FC<InputTextProps> = ({ id,label, value, onChange, size = 'm', minLength, maxLength, validate }) => {
    const [error, setError] = useState<string | null>(null);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        const newValue = event.target.value;

        onChange(newValue);
        setError(null);

        if (validate === false) {
            return;
        }

        // Validación de longitud
        if ((minLength && newValue.length < minLength) || (maxLength && newValue.length > maxLength)) {
            setError(`El texto debe tener entre ${minLength ?? '-' } y ${maxLength ?? '-' } caracteres`);
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
            error={!!error}
            helperText={error || ' '}
            fullWidth
        />
    );
};

export default InputText;
