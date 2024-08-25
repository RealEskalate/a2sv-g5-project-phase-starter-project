  import 'package:flutter/material.dart';

  import '../widget/recieved_messages.dart';
  import '../widget/sent_message.dart';
  import '../widget/user_profile.dart';

  class IndividiualChatScreen extends StatefulWidget {
  const IndividiualChatScreen({super.key});

  @override
  State<IndividiualChatScreen> createState() => _IndividiualChatScreenState();
}

class _IndividiualChatScreenState extends State<IndividiualChatScreen> {
  bool showSendIcon = false;
  TextEditingController searchTextController = TextEditingController();
  @override
    Widget build(BuildContext context) {
      final dummyMessage = [
        showRecievedMessage('This is my new 3d design'),
        showSentMessage('You did your job well!'),
        showRecievedMessage('This is my new 3d design'),
        showSentMessage('You did your job well!'),
        showRecievedMessage('This is my new 3d design'),
        showSentMessage('You did your job well!'),
        showRecievedMessage('This is my new 3d design'),
        showSentMessage('You did your job well!'),
        showRecievedMessage('This is my new 3d design'),
        showSentMessage('You did your job well!'),
        
      ];

      return Scaffold(
        
            appBar: AppBar(
              toolbarHeight: 70,
              centerTitle: false,
              automaticallyImplyLeading: false,
              flexibleSpace:const Padding(
                padding: EdgeInsets.only(top: 15.0),
                child: MyAppBar(),
              ),
              actions: [
                IconButton(onPressed: (){}, icon: const Icon(Icons.phone)),
              IconButton(onPressed: (){}, icon: const Icon(Icons.video_call))

              ],

        ),
        body: Column(
          children: [
            Expanded(
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: ListView.builder(
                  itemCount: 10,
                  shrinkWrap: true,
                  itemBuilder: (context, index) => dummyMessage[index]
                  ),
              ),
            ),
             InputBar(
              showSendbuttom: showSendIcon,
              textController: searchTextController,
              onChanged: (message) {
                if(message==''){
                  setState(() {
                    showSendIcon = false;
                  });
                }
                else{
                  setState(() {
                    showSendIcon = true;
                  });
                }
              },
            )
        
          ],
        ),
      );
    }
}

class MyAppBar extends StatelessWidget {
  const MyAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Row(
        children: [
        IconButton(onPressed: (){}, icon: const Icon(Icons.arrow_back)),
        showUser(),
        const Column(
          children: [
            Text('Sabila Sayima',style: TextStyle(fontWeight: FontWeight.bold,fontSize: 18),),
            Text('online')
          ],
        )
        ],
      ),
    );
  }
}



  class InputBar extends StatelessWidget {
    const InputBar({
      super.key, required this.textController, required this.onChanged, required this.showSendbuttom,
    });

    final TextEditingController textController;
    final void Function(String message) onChanged;
    final bool showSendbuttom;

    @override
    Widget build(BuildContext context) {
      return Padding(
              padding: const EdgeInsets.only(right: 8.0, bottom: 25.0),
              child: Row(
      children: [
        IconButton(
          onPressed: () {},
          icon: const Icon(Icons.attachment_sharp, color: Colors.blue),
        ),
        Expanded(
          child: TextField(
            onChanged: (value) {
              onChanged(value);
            },
            controller: textController,
            decoration: InputDecoration(
              suffixIcon: showSendbuttom?IconButton(onPressed: (){}, icon: const Icon(Icons.send)):null,
              hintText: 'Type a message...',
              contentPadding: const EdgeInsets.symmetric(horizontal: 10),
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(25),
                borderSide: const BorderSide(color: Colors.blue),
              ),
            ),
          ),
        ),
        
          IconButton(
          onPressed: () {},
          icon: const Icon(Icons.camera_alt, color: Colors.blue),
        ),
        IconButton(
          onPressed: () {},
          icon: const Icon(Icons.mic, color: Colors.blue),
        ),
      ],
              ),
            );
    }
  }