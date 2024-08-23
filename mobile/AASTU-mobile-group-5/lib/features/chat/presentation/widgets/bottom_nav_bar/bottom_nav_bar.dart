import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:google_fonts/google_fonts.dart';
import 'file_picker_widget.dart';
import 'image_capture_widget.dart';
import 'voice_recording_widget.dart';

class CustomBottomNavigationBar extends StatefulWidget {
  @override
  _CustomBottomNavigationBarState createState() =>
      _CustomBottomNavigationBarState();
}

class _CustomBottomNavigationBarState extends State<CustomBottomNavigationBar> {
  TextEditingController _controller = TextEditingController();
  bool _isTextFieldEmpty = true;
  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return BottomAppBar(
      color: Colors.white,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: <Widget>[
          // IconButton(
          //   icon: Icon(Icons.attach_file),
          //   onPressed: () {
          //     // Handle file upload
          //   },
          // ),
          Transform.rotate(
            angle: 0.5, 
            child: FilePickerWidget(),
          ),
          Expanded(
            child: TextField(
              decoration: InputDecoration(
                hintText: 'Write your message...',
                hintStyle: TextStyle(
                  // fontFamily: 'General Sans Variable',
                  fontWeight: FontWeight.w400,
                  fontSize: 14,
                  height: 1, // Line height is 12px, same as font size
                  color: Color.fromRGBO(121, 124, 123, 1),
                ),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(15),
                  borderSide: BorderSide.none,
                ),
                contentPadding: EdgeInsets.symmetric(horizontal: 10),
                filled: true,
                fillColor: Colors.grey[200],
                focusedBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(30),
                  borderSide: BorderSide.none,
                ),
                suffixIcon: IconButton(
                  icon: Icon(Icons.copy_all_rounded),
                  onPressed: () {
                    // Handle copy action
                    Clipboard.setData(ClipboardData(text: _controller.text));
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(content: Text('Copied to clipboard')),
                    );
                  },
                ),
              ),
              controller: _controller,
              onChanged: (text) {
                setState(() {
                  _isTextFieldEmpty = text.isEmpty;
                });
              },
            ),
          ),
          // IconButton(
          //   icon: Icon(Icons.camera_alt),
          //   onPressed: () {
          //     // Handle image picker
          //   },
          // ),
          Visibility(visible: _isTextFieldEmpty, child: ImageCaptureWidget()),
          // IconButton(
          //   icon: Icon(Icons.mic),
          //   onPressed: () {
          //     // Handle voice recording
          //   },
          // ),
          Visibility(visible: _isTextFieldEmpty, child: VoiceRecordingWidget()),
          if (!_isTextFieldEmpty)
            IconButton(
              icon: Icon(Icons.send),
              onPressed: () {
                // Handle send action
              },
            ),
        ],
      ),
    );
  }
}
