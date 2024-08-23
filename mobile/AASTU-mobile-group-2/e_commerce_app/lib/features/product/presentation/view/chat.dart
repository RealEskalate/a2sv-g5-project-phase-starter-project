import 'package:e_commerce_app/features/auth/presentation/view/widgets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_sound/flutter_sound.dart';
import 'package:image_picker/image_picker.dart';
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'dart:io';
import 'package:chat_bubbles/bubbles/bubble_normal_audio.dart';
import 'package:bubble/bubble.dart';

class Chat extends StatefulWidget {
  const Chat({super.key});

  @override
  _ChatState createState() => _ChatState();
}

class _ChatState extends State<Chat> {
  final TextEditingController _controller = TextEditingController();
   final ScrollController _scrollController = ScrollController();
  final List<Message> messages = [
    Message(
      type: 'text',
      content: 'Hello!',
      isUser: true,
      name: 'User',
      date: '8/23/2024',
      profileImage: File('assets/smile.png'),
    ),
    Message(
      type: 'image',
      content: 'assets/smile.png',
      isUser: false,
      name: 'Bot',
      date: '8/23/2024',
      profileImage: File('assets/smile.png'),
    ),
    // Message(
    //   type: 'voice',
    //   content: 'voice.mp3',
    //   isUser: true,
    //   name: 'User',
    //   date: '8/23/2024',
    //   profileImage: File('images/image.png'),
    // ),
  ];

  // Initializing audio-related variables'
  final AudioUtils _audioUtils = AudioUtils();
  Duration position = Duration();
  bool isPlaying = false;
  bool isLoading = false;
  bool isPause = false;
  // final Timer ? _timer;
  FlutterSoundRecorder? _recorder;
  // FlutterSoundPlayer? _player;
  String? _filePath;
  bool _isRecording = false;
  bool _isPlaying = false;
  FlutterSoundPlayer? _player;

  String? selectedImage;

  //
  final ImagePicker picker = ImagePicker();

  Future pickImage() async {
    final pickedImage =
        await ImagePicker().pickImage(source: ImageSource.gallery);
    if (pickedImage == null) return;
    setState(() {
      messages.add(Message(
          type: "image",
          content: pickedImage.path,
          isUser: true,
          name: "name",
          date: "date"));
      selectedImage = pickedImage.path;
    });
    _scroll();
  }
  void _scroll() {
     FocusScope.of(context).unfocus();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (_scrollController.hasClients) {
        _scrollController.animateTo(
          _scrollController.position.maxScrollExtent,
          duration: Duration(milliseconds: 200),
          curve: Curves.easeOut,
        );
      }
    });
  }
// Duration _getDuration(String filePath) async{
//    return await _audioUtils.getAudioDuration(filePath) ?? Duration.zero;
//   }
  @override
  void initState() {
    super.initState();
    _recorder = FlutterSoundRecorder();
    _player = FlutterSoundPlayer();
    _initRecorder();
  }

  Future<void> _initRecorder() async {
    await _recorder!.openRecorder();
    await _player!.openPlayer();
    await _requestPermissions();
  }

  Future<void> _requestPermissions() async {
    await Permission.microphone.request();
    await Permission.storage.request();
  }

  Future<void> _startRecording() async {
    print("start rec");
    final directory = await getApplicationDocumentsDirectory();
    _filePath = '${directory.path}/audio_example.aac';
    print(_filePath);

    await _recorder!.startRecorder(
      toFile: _filePath,
      codec: Codec.aacADTS,
    );

    setState(() {
      _isRecording = true;
      // _timer = Timer.periodic(Duration(seconds: 1), (Timer timer) {
      //   setState(() {
      //     _elapsedDuration += Duration(seconds: 1);
      //   });
      // });
    });
  }

  Future<void> _stopRecording() async {
    await _recorder!.stopRecorder();
    messages.add(Message(
        type: "voice",
        content: _filePath!,
        isUser: true,
        name: "trial voice",
        date: "date"));

    // print(getFileDuration(_filePath!));
    _scroll();
    setState(() {
      _isRecording = false;
    
    });
  }

  Future<void> _playAudio() async {
    if (_filePath != null && File(_filePath!).existsSync()) {
      await _player!.startPlayer(
        fromURI: _filePath,
        codec: Codec.aacADTS,
        whenFinished: () {
          setState(() {
            _isPlaying = false;
          });
        },
      );

      setState(() {
        _isPlaying = true;
      });
    }
  }

  Future<void> _stopAudio() async {
    await _player!.stopPlayer();
    setState(() {
      _isPlaying = false;
    });
  }

  @override
  void dispose() {
    _recorder!.closeRecorder();
    _player!.closePlayer();
    super.dispose();
  }

  void _sendMessage() {
    if (_controller.text.isNotEmpty) {
      // FocusScope.of(context).unfocus();
      setState(() {
        messages.add(Message(
            type: "text",
            content: _controller.text,
            isUser: true,
            name: "name",
            date: "date"));
        _controller.clear();
      });
      _scroll();
    }
  }

  void _sendImage() {
    // Implement image sending functionality
  }

  void _recordAudion() {
    // Implement audio sending functionality
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    double screenHeight = MediaQuery.of(context).size.height;

    return Scaffold(
      appBar: AppBar(
        title: Row(
          children: [
            IconButton(
              onPressed: () {},
              icon: const Icon(Icons.arrow_back),
            ),
            const SizedBox(width: 8),
            Container(
              width: screenWidth * 0.15,
              height: screenWidth * 0.15,
              child: Stack(
                alignment: Alignment.center,
                children: [
                  Icon(
                    Icons.circle,
                    color: Color(0xFFFEC7D3),
                    size: screenWidth * 0.15,
                  ),
                  Image(
                    image: AssetImage('assets/smile.png'),
                    width: screenWidth * 0.12,
                    height: screenWidth * 0.12,
                  ),
                  Positioned(
                    bottom: 5,
                    right: 5,
                    child: Icon(
                      Icons.circle,
                      color: Colors.green,
                      size: screenWidth * 0.03,
                    ),
                  ),
                ],
              ),
            ),
            SizedBox(width: 8),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    'Sabila Sayma',
                    style: TextStyle(
                      fontSize: screenWidth * 0.05,
                    ),
                  ),
                  Text(
                    'online',
                    style: TextStyle(
                      fontSize: screenWidth * 0.03,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ],
              ),
            ),
            SizedBox(width: 16),
            Icon(Icons.call, size: screenWidth * 0.07),
            SizedBox(width: 16),
            Icon(Icons.video_call, size: screenWidth * 0.07),
          ],
        ),
      ),
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.end,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          // mebebt
          // Expanded(
          //   child: ListView.builder(
          //     itemCount: _messages.length,
          //     itemBuilder: (context, index) {
          //       return ListTile(
          //         title: Text(_messages[index]),
          //       );
          //     },
          //   ),
          // ),

          Expanded(
            child: ListView.builder(
              controller: _scrollController,
              itemCount: messages.length,
              itemBuilder: (context, index) {
                final message = messages[index];
                return Padding(
                  padding:
                      EdgeInsets.symmetric(vertical: 8.0, horizontal: 16.0),
                  child: Row(
                    mainAxisAlignment: message.isUser
                        ? MainAxisAlignment.start
                        : MainAxisAlignment.end,
                    children: [
                      if (message.isUser)
                        _buildProfileAvatar(message.profileImage),
                      SizedBox(width: 10),
                      Expanded(child: _buildMessageContent(message)),
                      if (!message.isUser)
                        _buildProfileAvatar(message.profileImage),
                    ],
                  ),
                );
              },
            ),
          ),

          ///mjd
          // Container(
          //     child: selectedImage != null
          //         ? Image.file(
          //             fit: BoxFit.cover,
          //             File(selectedImage!),
          //           )
          //         : Container()),
          // if (_filePath != null)
          //   ElevatedButton(
          //     onPressed: _isPlaying ? _stopAudio : _playAudio,
          //     child: Text(_isPlaying ? 'Stop Playback' : 'Play Audio'),
          //   ),
          // SizedBox(height: 20),
          // if (_filePath != null) Text('Recorded file: $_filePath'),

          //debug
          Container(
            color: Colors.white,
            padding: const EdgeInsets.only(top: 10, bottom: 10),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                IconButton(
                  icon: Icon(Icons.attach_file),
                  onPressed: null,
                ),
                Expanded(
                  child: TextFieldTitle(
                    controller: _controller,
                    hint: 'Type a message',
                    iconButton: IconButton(
                        icon: Icon(Icons.send), onPressed: _sendMessage),
                  ),
                ),
                IconButton(
                    icon: Icon(Icons.camera_alt),
                    onPressed: () {
                      pickImage();
                    }),
                _isRecording ? Text("recor") : Container(),
                IconButton(
                  icon: _isRecording ? Icon(Icons.stop) : Icon(Icons.mic),
                  onPressed: () {
                    _isRecording ? _stopRecording() : _startRecording();
                  },
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  ///mer

  Widget _buildProfileAvatar(File? profileImage) {
    return CircleAvatar(
      backgroundImage: profileImage != null ? FileImage(profileImage) : null,
      child: profileImage == null ? Icon(Icons.person) : null,
    );
  }

  Widget _buildMessageContent(Message message) {
    
    // duration = await FlutterSoundHelper().duration(File(message.content));
    return Column(
      crossAxisAlignment:
          message.isUser ? CrossAxisAlignment.start : CrossAxisAlignment.end,
      children: [
        Text(
          message.name,
          style: TextStyle(fontWeight: FontWeight.bold, fontSize: 14),
        ),
        SizedBox(height: 5),
        _buildMessageBubble(message),
        SizedBox(height: 5),
        Text(
          message.date,
          style: TextStyle(fontSize: 12, color: Colors.grey),
        ),
      ],
    );
  }

  Widget _buildMessageBubble(Message message) {
     
    switch (message.type) {
      case 'text':
        return Bubble(
          child: Text(message.content),
          color: message.isUser ? Colors.blue : Colors.grey[300],
          nip: message.isUser ? BubbleNip.leftTop : BubbleNip.rightTop,
        );
      case 'image':
        return Container(
          constraints: BoxConstraints(maxWidth: 200),
          child: Image.file(File(message.content)),
        );
      case 'voice':
      // stop here
        return BubbleNormalAudio(
          // constraints: BoxConstraints(maxWidth: screenWidth * 0.75),
          color: Color(0xFFE8E8EE),
          duration: 60,
          position: 7,
          isPlaying: _isPlaying,
          isLoading: isLoading,
          isPause: _isPlaying ? false : true,
          onSeekChanged: _changeSeek,
          onPlayPauseButtonClick: () {
            print("play pause");
            _isPlaying ? _stopAudio() : _playAudio();
          },
          sent: message.isUser,
          
        );
      default:
        return SizedBox.shrink();
    }
  }

  void _changeSeek(double value) {
    position = Duration(seconds: value.toInt());
  }

  // void _playAudio() {
  //   isPlaying = !isPlaying;
  // }
}

class Message {
  final String type; // 'text', 'image', or 'voice'
  final String content; // message text or image/voice file path
  final bool isUser;
  final String name;
  final String date;
  final File? profileImage;

  Message({
    required this.type,
    required this.content,
    required this.isUser,
    required this.name,
    required this.date,
    this.profileImage,
  });
}

class AudioUtils {
  final FlutterSoundPlayer _player = FlutterSoundPlayer();

  Future<Duration?> getAudioDuration(String filePath) async {
    try {
      await _player.openPlayer();
      Duration? duration = await _player.startPlayer(
        fromURI: filePath,
        whenFinished: () {
          _player.stopPlayer();
        },
      );
      await _player.stopPlayer();
      return duration;
    } catch (e) {
      print("Error retrieving audio duration: $e");
      return null;
    } finally {
      await _player.closePlayer();
    }
  }
}
