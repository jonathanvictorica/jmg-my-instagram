import React, { useState, useEffect } from 'react';
import { TextField, Autocomplete } from '@mui/material';
import { styled } from '@mui/material/styles';

interface AutocompleteInputProps {
    id: string;
    label: string;
    value: string;
    onChange: (value: string) => void;
    fetchData: (query: string) => Promise<string[]>; // Función para obtener datos del backend
    size?: 's' | 'm' | 'l'| 'xl'| 'xxl';
}

const StyledAutocomplete = styled(Autocomplete)<{ inputSize: string }>(({ inputSize }) => {
    let inputWidth: string;

    inputWidth = getSize(inputSize);

    return {
        margin: 5,
        '& .MuiInputBase-root': {
            minWidth: inputWidth,
            maxWidth: inputWidth,
        },
    };
});

const InputAutocomplete: React.FC<AutocompleteInputProps> = ({
    id,
                                                                 label,
                                                                 value,
                                                                 onChange,
                                                                 fetchData,
                                                                 size = 'm',
                                                             }) => {
    const [options, setOptions] = useState<string[]>([]); // Opciones de autocompletado
    const [loading, setLoading] = useState<boolean>(false); // Indicador de carga

    // useEffect para manejar la carga de opciones en base a la entrada del usuario
    useEffect(() => {
        if (value.length > 2) { // Solo buscar cuando el valor tenga más de 2 caracteres
            setLoading(true);
            fetchData(value)  // Hacer la solicitud al backend
                .then((data) => {
                    setOptions(data);
                    setLoading(false);
                })
                .catch(() => {
                    setOptions([]);
                    setLoading(false);
                });
        } else {
            setOptions([]);
        }
    }, [value, fetchData]); // Se ejecuta cuando cambia el valor del input

    return (
        <StyledAutocomplete
            id={id}
            inputSize={size} // Propiedad para ajustar el tamaño del input
            freeSolo
            options={options} // Opciones de autocompletado
            loading={loading} // Muestra el indicador de carga
            onInputChange={(_, newInputValue) => onChange(newInputValue)} // Actualiza el valor al escribir
            renderInput={(params) => (
                <TextField
                    {...params}
                    label={label}
                    variant="outlined"
                    size={size}
                    fullWidth
                    InputProps={{
                        ...params.InputProps,
                        endAdornment: loading ? <div>Loading...</div> : null, // Indicador de carga al final del input
                    }}
                />
            )}
        />
    );
};

export default InputAutocomplete;
