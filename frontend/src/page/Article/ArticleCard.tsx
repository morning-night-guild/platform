import { Box, Heading, Stack, useColorModeValue, Image } from '@chakra-ui/react';

export type ArticleCardProps = {
    title: string;
    thumbnailURL?: string;
    url: string;
};

export function ArticleCard(props: ArticleCardProps) {
    return (
        <Box w="full" rounded="md" p={6} overflow="hidden">
            <Stack>
                <Box as="a" href={props.url} target="_blank">
                    <Image src={props.thumbnailURL ?? './noimage.png'} width="100%" fallbackSrc="./noimage.png" />
                </Box>
                <Heading color={useColorModeValue('gray.700', 'white')} fontSize="xl" fontFamily="body">
                    {props.title}
                </Heading>
            </Stack>
        </Box>
    );
}
