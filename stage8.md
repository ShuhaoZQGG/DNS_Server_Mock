1. Tester => (DNS Packet 1) => DNS Server (me) => ([] DNS Packet 2) => Resolver

In step1, the DNS packet 1 contains header, []questions
Seperate []questions into individual question, 
DNS Packet 2 slices should each contain header, question

2. Resolver => (DNS Packet 3) => DNS Server (me) => (DNS Packet 4) => Tester

DNS Packet 3 will contain a header,  a question, and an answer
DNS Packet 4 will be the same as DNS Packet 3



In code:
1. check if received dns packet have multiple questions, if so, reconstruct it into multiple dns packets with only one question
2. if received dns packet has only one question and also contains answer, do nothing but direcly send the dns packets


2 sources: one tester, another resolver

how to tell which is which:
tester will have dns packet with header with ANCOUNT = 0
resolver will have dns packet with header with ANCOUNT > 0

tester will have data after parsing like this: DNS { Header, []Questions }
we need to send to resolvers data after parsing lke this : [] DNS { Header, Question}

resolver will have data after parsing like this: [] DNS { Header, Question, Answer}
we need to send to tester data after parsing like this: DNS { Header, []Questions, []Answers }