import { Flex, Divider, Heading, Image, Center, Text } from '@chakra-ui/react';

export type HeaderProps = Record<string, unknown>;

export function Header(props: HeaderProps) {
    return (
        <>
            <Center>
                <Flex as="nav" align="center" wrap="wrap" pt={3} pb={3} width="80%">
                    <Heading as="h1" size="lg" letterSpacing="tighter" _hover={{ color: 'brand.main' }}>
                        <Center>
                            <Image src="/logo.png" boxSize="40px" />
                            <Text display={['none', 'inline', 'inline', 'inline']}>Platform</Text>
                        </Center>
                    </Heading>
                </Flex>
            </Center>
            <Divider />
        </>
    );
}
