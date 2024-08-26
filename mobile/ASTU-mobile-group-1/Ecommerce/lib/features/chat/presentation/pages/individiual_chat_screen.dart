import 'package:flutter/material.dart';

import '../../../auth/presentation/pages/pages.dart';
import '../bloc/chat_bloc.dart';
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
  TextEditingController sendTextController = TextEditingController();
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
        flexibleSpace: Padding(
          padding: const EdgeInsets.only(top: 15.0),
          child: MyAppBar(
            onBack: () => Navigator.of(context).pop(),
          ),
        ),
        actions: [
          IconButton(onPressed: () {}, icon: const Icon(Icons.phone)),
          IconButton(onPressed: () {}, icon: const Icon(Icons.video_call))
        ],
      ),
      body: BlocConsumer<ChatBloc, ChatState>(
        listener: (context, state) {
           if(state is ChatFailureState){
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.message))
            );
           }
        },
        builder: (context, state) {
          return Column(
            children: [
              Expanded(
                child: Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: ListView.builder(
                      itemCount: 10,
                      shrinkWrap: true,
                      itemBuilder: (context, index) => dummyMessage[index]),
                ),
              ),
              InputBar(
                onSend: (){
                  if(state is IndividualChatingState){
                  context.read<ChatBloc>().add(SendMessageEvent(message: sendTextController.text, type: 'text', chatEntity: state.chatEntity ));
                  }
                },
                showSendbuttom: showSendIcon,
                textController: sendTextController,
                onChanged: (message) {
                  if (message == '') {
                    setState(() {
                      showSendIcon = false;
                    });
                  } else {
                    setState(() {
                      showSendIcon = true;
                    });
                  }
                },
              )
            ],
          );
        },
      ),
    );
  }
}

class MyAppBar extends StatelessWidget {
  const MyAppBar({super.key, required this.onBack});
  final VoidCallback onBack;
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Row(
        children: [
          IconButton(
              onPressed: () {
                onBack();
              },
              icon: const Icon(Icons.arrow_back)),
          showUser(onClicked: () {}),
          const Column(
            children: [
              Text(
                'Sabila Sayima',
                style: TextStyle(fontWeight: FontWeight.bold, fontSize: 18),
              ),
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
    super.key,
    required this.textController,
    required this.onChanged,
    required this.showSendbuttom, required this.onSend,
  });

  final TextEditingController textController;
  final void Function(String message) onChanged;
  final bool showSendbuttom;
  final VoidCallback onSend;


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
                suffixIcon: showSendbuttom
                    ? IconButton(onPressed: () {onSend();}, icon: const Icon(Icons.send))
                    : null,
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
