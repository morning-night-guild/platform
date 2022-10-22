import { Grid, GridItem, Center, Heading, Box } from '@chakra-ui/react';
import { Header } from '../Header';
import { ArticleCard } from './ArticleCard';

export function Article() {
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
                            <GridItem colSpan={1}>
                                <ArticleCard
                                    thumbnailURL="https://d2908q01vomqb2.cloudfront.net/887309d048beef83ad3eabf2a79a64a389ab1c9f/2022/09/26/DBBLOG-2257-featured-images.jpg"
                                    url="https://aws.amazon.com/blogs/database/build-a-cqrs-event-store-with-amazon-dynamodb/"
                                    title="Build a CQRS event store with Amazon DynamoDB"
                                />
                            </GridItem>
                            <GridItem colSpan={1}>
                                <ArticleCard
                                    thumbnailURL="https://cdn-ak.f.st-hatena.com/images/fotolife/v/vasilyjp/20220520/20220520131721.jpg"
                                    url="https://techblog.zozo.com/entry/improve-digdag-task-with-goroutine"
                                    title="Goプログラム実行時間の短縮 - ZOZO TECH BLOG"
                                />
                            </GridItem>
                            <GridItem colSpan={1}>
                                <ArticleCard
                                    thumbnailURL="https://example.com"
                                    url="https://github.com/EbookFoundation/free-programming-books/blob/main/books/free-programming-books-ja.md"
                                    title="free-programming-books/free-programming-books-ja.md at main · EbookFoundation/free-programming-books"
                                />
                            </GridItem>
                            <GridItem colSpan={1}>
                                <ArticleCard
                                    thumbnailURL="https://res.cloudinary.com/zenn/image/upload/s--Ib8D-lJf--/co_rgb:222%2Cg_south_west%2Cl_text:notosansjp-medium.otf_37_bold:%25E3%2581%2582%25E3%2581%2584%25E3%2581%2597%25E3%2583%25BC%2Cx_203%2Cy_98/c_fit%2Cco_rgb:222%2Cg_north_west%2Cl_text:notosansjp-medium.otf_70_bold:Hugo%25E3%2581%25A7%25E3%2581%258A%25E6%2589%258B%25E8%25BB%25BD%25E6%259C%2580%25E5%25BC%25B7%25E3%2583%259D%25E3%2583%25BC%25E3%2583%2588%25E3%2583%2595%25E3%2582%25A9%25E3%2583%25AA%25E3%2582%25AA%25E3%2582%2592%25E4%25BD%259C%25E3%2582%258B%2Cw_1010%2Cx_90%2Cy_100/g_south_west%2Ch_90%2Cl_fetch:aHR0cHM6Ly9yZXMuY2xvdWRpbmFyeS5jb20vemVubi9pbWFnZS9mZXRjaC9zLS1QQVVSWmo3US0tL2NfbGltaXQlMkNmX2F1dG8lMkNmbF9wcm9ncmVzc2l2ZSUyQ3FfYXV0byUyQ3dfNzAvaHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL3plbm4tdXNlci11cGxvYWQvYXZhdGFyLzI1NzIyOGFiMWEuanBlZw==%2Cr_max%2Cw_90%2Cx_87%2Cy_72/v1627274783/default/og-base_z4sxah.png"
                                    url="https://zenn.dev/1see/articles/093da70d4dada1"
                                    title="Hugoでお手軽最強ポートフォリオを作る"
                                />
                            </GridItem>
                        </Grid>
                    </Center>
                </GridItem>
                <GridItem colSpan={2} />
            </Grid>
        </>
    );
}
