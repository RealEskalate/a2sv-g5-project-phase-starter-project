import 'package:ecommerce_app/features/chat/presentation/pages/chat_room.dart';
import 'package:flutter/material.dart';

class ChatList extends StatefulWidget {
  const ChatList({super.key});

  @override
  State<ChatList> createState() => _ChatListState();
}

class _ChatListState extends State<ChatList> {
  List imagesList = [
    'assets/images/Memoji Boys 6-18.png',
    'assets/images/42 11.png',
    'assets/images/21 5.png'
  ];
  List nameList = ['My status', 'Adil', 'Marina', 'Dean', 'Max'];

  List<Color> colorList = [
    const Color(0xffC5C5C5),
    const Color(0xffFBE8A8),
    const Color(0xffFEC7D3),
    const Color(0xffC8ECFD),
    const Color(0xffC7FEE0)
  ];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: const Color(0xff498CF0),
        body: SafeArea(
          child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Padding(
                  padding: EdgeInsets.only(left: 10.0, top: 15, bottom: 20),
                  child: Icon(Icons.search),
                ),
                Padding(
                  padding:
                      const EdgeInsets.symmetric(horizontal: 8.0, vertical: 12),
                  child: SizedBox(
                    height: 89,
                    child: ListView.builder(
                      itemCount: 8,
                      scrollDirection: Axis.horizontal,
                      itemBuilder: (context, index) {
                        return Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 8.0),
                          child: Column(
                            children: [
                              Stack(
                                alignment: Alignment.center,
                                children: [
                                  Container(
                                    height: 60,
                                    width: 60,
                                    decoration: BoxDecoration(
                                      shape: BoxShape.circle,
                                      border: Border.all(
                                        color: colorList[index % 5],
                                        width: 3.0,
                                      ),
                                    ),
                                  ),
                                  Container(
                                    height: 50,
                                    width: 50,
                                    decoration: BoxDecoration(
                                      color: colorList[index % 5],
                                      shape: BoxShape.circle,
                                    ),
                                    child: ClipRRect(
                                      borderRadius: BorderRadius.circular(50),
                                      child: Image.asset(imagesList[index % 3]),
                                    ),
                                  ),
                                ],
                              ),
                              Text(nameList[index % 5])
                            ],
                          ),
                        );
                      },
                    ),
                  ),
                ),
                Expanded(
                  child: DecoratedBox(
                      decoration: const BoxDecoration(
                          color: Colors.white,
                          borderRadius:
                              BorderRadius.vertical(top: Radius.circular(40))),
                      child: Padding(
                        padding: const EdgeInsets.only(top: 8.0),
                        child: ListView.builder(
                          itemCount: 8,
                          padding: const EdgeInsets.all(8),
                          scrollDirection: Axis.vertical,
                          itemBuilder: (context, index) {
                            return Padding(
                              padding: const EdgeInsets.symmetric(
                                  vertical: 8.0, horizontal: 15),
                              child: GestureDetector(
                                onTap: () {
                                  Navigator.push(
                                      context,
                                      MaterialPageRoute(
                                        builder: (context) => const ChatRoom(),
                                      ));
                                },
                                child: Row(
                                  children: [
                                    Container(
                                      height: 50,
                                      width: 50,
                                      decoration: BoxDecoration(
                                        color: colorList[index % 5],
                                        shape: BoxShape.circle,
                                      ),
                                      child: ClipRRect(
                                        borderRadius: BorderRadius.circular(50),
                                        child:
                                            Image.asset(imagesList[index % 3]),
                                      ),
                                    ),
                                    const Padding(
                                      padding: EdgeInsets.only(left: 8.0),
                                      child: Column(
                                        children: [
                                          Text(
                                            'Alex Linderson',
                                            style: TextStyle(
                                                fontWeight: FontWeight.w600,
                                                fontSize: 15),
                                          ),
                                          Text(
                                            'How are you today?',
                                            style: TextStyle(
                                                fontWeight: FontWeight.w400,
                                                fontSize: 12,
                                                color: Colors.grey),
                                          ),
                                        ],
                                      ),
                                    ),
                                    const Spacer(),
                                    const Column(
                                      children: [
                                        Text(
                                          '2 min ago',
                                          style: TextStyle(
                                              fontWeight: FontWeight.w400,
                                              fontSize: 12,
                                              color: Colors.grey),
                                        ),
                                        DecoratedBox(
                                          decoration: BoxDecoration(
                                              shape: BoxShape.circle,
                                              color: Color(0xff498CF0)),
                                          child: Padding(
                                            padding: EdgeInsets.all(8.0),
                                            child: Center(
                                                child: Text(
                                              '3',
                                              style: TextStyle(
                                                  color: Colors.white),
                                            )),
                                          ),
                                        )
                                      ],
                                    ),
                                  ],
                                ),
                              ),
                            );
                          },
                        ),
                      )),
                )
              ]),
        ));
  }
}
