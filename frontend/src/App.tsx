import { ChakraProvider } from '@chakra-ui/react';
import { Article } from './page/Article';
import { theme } from './theme';

function App() {
    return (
        <ChakraProvider theme={theme}>
            <Article />
        </ChakraProvider>
    );
}

export default App;
