import React, {useState} from 'react';
import ButtonAtom from "../components/atom/Button";
import InputNumber from "../components/atom/InputNumber";
import InputText from "../components/atom/InputText";
import InputAutocomplete from "../components/atom/InputTextAutocomplet";
import InputSelect from "../components/atom/InputSelect";

const Home: React.FC = () => {
    const [selectedOption, setSelectedOption] = useState<string>('');
    const [age,setAge]= useState(0)

    const [validate, setValidate] = useState(false);
    const [nameAutocomplete, setNameAutocomplete] = useState<string>('');
    const [name,setName]= useState('')

    const fetchDataOpciones = (): Promise<{ key: string; value: string }[]> => {
        return new Promise((resolve) => {
            setTimeout(() => {
                const options = [
                    { key: '1', value: 'Opción 1' },
                    { key: '2', value: 'Opción 2' },
                    { key: '3', value: 'Opción 3' },
                    { key: '4', value: 'Opción 4' },
                ];
                resolve(options);
            }, 10); // Simulamos un retraso de 1 segundo
        });
    };


    const fetchData = (query: string): Promise<string[]> => {
        return new Promise((resolve) => {
            // Simulamos un retraso de 1 segundo (1000 ms)
            setTimeout(() => {
                // Generamos una lista de opciones simuladas basadas en el texto que ingresa el usuario
                const data = [
                    'Juan',
                    'Juliana',
                    'José',
                    'Javier',
                    'Julia',
                    'Jessica',
                    'Joaquín',
                    'Jimmy',
                    'José Luis',
                    'Jimena',
                ];

                // Filtramos las opciones que contienen la consulta 'query'
                const filteredData = data.filter((item) =>
                    item.toLowerCase().includes(query.toLowerCase())
                );

                resolve(filteredData); // Resolvemos la promesa con las opciones filtradas
            }, 1000); // Simulamos un retraso de 1 segundo
        });
    };






    return (
        <>
            <h2>Home Page</h2>
            <p>Welcome to the homepage!</p>

            <InputAutocomplete
                id={"444"}
                label="Nombre Autocomplete"
                value={nameAutocomplete}
                onChange={setNameAutocomplete}
                fetchData={fetchData}
                size="m"
            />
            <InputText
                id={"222"}
                validate={validate}
                label="Nombre"
                value={name}
                onChange={(newValue) => setName(newValue)}
                size="m"
                minLength={3}
                maxLength={50}
            />

            <InputNumber
                id={"111"}
                validate={validate}
                label="Edad"
                value={age}
                onChange={(newValue) => setAge(newValue)}
                size="m"
                min={18}
                max={99}
            />

            <ButtonAtom id={"1"} label="Guardar" type="save" size="m" onClick={() => setValidate(!validate)} />
            <ButtonAtom id={"2"} label="Cancelar" type="cancel" size="s" onClick={() => setValidate(!validate)} />
            <ButtonAtom id={"3"} label="Continuar" type="continue" size="l" onClick={() => setValidate(!validate)} />


        </>
    );
};

export default Home;
