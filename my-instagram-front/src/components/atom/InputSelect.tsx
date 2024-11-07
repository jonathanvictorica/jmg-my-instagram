import React, { useState, useEffect } from 'react';
import { Select, MenuItem, InputLabel, FormControl } from '@mui/material';
import { styled } from '@mui/material/styles';

interface SelectOption {
    key: string;
    value: string;
}

interface InputSelectProps {
    id: string;
    label: string;
    value: string;
    onChange: (value: string) => void;
    fetchData: () => Promise<SelectOption[]>; // Función que devuelve las opciones con los campos key, value
    size: 's' | 'm' | 'l' | 'xl' | 'xxl';  // Tamaños posibles
}

// Función para obtener el tamaño basado en el valor de inputSize
function getSize(inputSize: 's' | 'm' | 'l' | 'xl' | 'xxl'): string {
    const sizeMap = {
        s: '80px',
        m: '120px',
        l: '160px',
        xl: '200px',
        xxl: '240px'
    };
    return sizeMap[inputSize] || '120px'; // Valor por defecto
}

const StyledFormControl = styled(FormControl)<{ inputSize: string }>(({ inputSize }) => {
    const inputWidth = getSize(inputSize); // Llamamos a getSize para obtener el tamaño

    return {
        margin: '5px', // Márgenes de 5px alrededor del componente
        '& .MuiInputBase-root': {
            minWidth: inputWidth,
            maxWidth: inputWidth,
        },
    };
});

const InputSelect: React.FC<InputSelectProps> = ({ id,label, value, onChange, fetchData, size }) => {
    const [options, setOptions] = useState<SelectOption[]>([]); // Opciones del select

    // useEffect para manejar la carga de opciones desde la función fetchData
    useEffect(() => {
        fetchData()  // Llama a la función fetchData que obtiene las opciones
            .then((data) => {
                setOptions(data);
            })
            .catch(() => {
                setOptions([]);
            });
    }, [fetchData]); // Solo se ejecuta una vez cuando el componente se monta

    return (
        <StyledFormControl inputSize={size} fullWidth variant="outlined">
            <InputLabel>{label}</InputLabel>
            <Select
                id={id}
                value={value}
                onChange={(event) => onChange(event.target.value)}
                label={label}
            >
                {options.map((option) => (
                    <MenuItem key={option.key} value={option.key}>
                        {option.value}
                    </MenuItem>
                ))}
            </Select>
        </StyledFormControl>
    );
};

export default InputSelect;
