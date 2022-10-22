import { Grid, GridItem, Center, Heading, Box } from '@chakra-ui/react';
import { Header } from '../Header';
import { useListArticles } from '../../swr/useApiArticle';
import { ArticleCard } from './ArticleCard';

export function Article() {
    const data = useListArticles();

    return (
        <>
            <Header />
            <Grid templateColumns="repeat(12, 1fr)">
                <GridItem colSpan={2} />
                <GridItem gap={4} colSpan={8}>
                    <Box m={4}>
                        <Heading>Articles</Heading>
                    </Box>
                    <Center>
                        <Grid
                            gap={6}
                            templateColumns={['repeat(1, 1fr)', 'repeat(1, 1fr)', 'repeat(2, 1fr)', 'repeat(2, 1fr)']}
                        >
                            {data?.articles?.map((article) => {
                                return (
                                    <GridItem key={article.id} colSpan={1}>
                                        <ArticleCard
                                            thumbnailURL={article.thumbnail}
                                            url={article.url}
                                            title={article.title}
                                        />
                                    </GridItem>
                                );
                            })}
                        </Grid>
                    </Center>
                </GridItem>
                <GridItem colSpan={2} />
            </Grid>
        </>
    );
}
