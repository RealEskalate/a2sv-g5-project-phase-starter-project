// Text(
//   'Your Text Here',
//   style: Theme.of(context).textTheme.bodyMedium?.copyWith(
//     fontSize: 18.0, // Your desired font size
//     fontWeight: FontWeight.bold, // Your desired font weight
//   ),
// ),



// Theme.of(context).colorScheme.onSurface,






// child: Row(
//                   // mainAxisAlignment: MainAxisAlignment.spaceBetween,
//                   children: [
//                     Row(
//                       mainAxisAlignment: MainAxisAlignment.spaceBetween,
//                       children: [
//                         Container(
//                           color: Colors.red,
//                           margin: EdgeInsets.only(top: 4),
//                           child: SizedBox(
//                               width: 80,
//                               height: 80,
//                               child: ImagePickerIconButton()),
//                         ),
//                         GestureDetector(
//                           onTap: () {
//                             Navigator.pushNamed(context, '/logout');
//                           },
//                           child: Container(
//                             padding: EdgeInsets.only(top: 4, left: 10),
//                             child: Column(
//                                 crossAxisAlignment: CrossAxisAlignment.start,
//                                 children: [
//                                   Text("July 31, 2024",
//                                       style: GoogleFonts.syne(
//                                         fontWeight: FontWeight.w500,
//                                         color: Theme.of(context)
//                                             .colorScheme
//                                             .onSurface,
//                                       )),
//                                   Row(children: [
//                                     Text("Hello,",
//                                         style: GoogleFonts.sora(
//                                             fontWeight: FontWeight.w400,
//                                             color: Color.fromARGB(
//                                                 255, 102, 102, 102))),
//                                     BlocBuilder<GetUserBloc, GetUserState>(
//                                       builder: (context, state) {
//                                         if (state is GetUserLoading) {
//                                           return Text("Fetching User...",
//                                               style: GoogleFonts.sora(
//                                                 fontWeight: FontWeight.w600,
//                                                 color: Theme.of(context)
//                                                     .colorScheme
//                                                     .onSurface,
//                                               ));
//                                         } else if (state is GetUserLoaded) {
//                                           return Text("${state.user.name}",
//                                               style: GoogleFonts.sora(
//                                                 fontWeight: FontWeight.w600,
//                                                 color: Theme.of(context)
//                                                     .colorScheme
//                                                     .onSurface,
//                                               ));
//                                         } else {
//                                           return Text('name');
//                                         }
//                                       },
//                                     ),
//                                   ])
//                                 ]),
//                           ),
//                         ),
//                       ],
//                     ),
//                     Row(
//                       mainAxisAlignment: MainAxisAlignment.end,
//                       children: [
//                         Container(
//                             decoration: BoxDecoration(
//                                 // border: Border.all(
//                                 // color: Color.fromRGBO(221, 221, 221, 1), width: 2),
//                                 borderRadius: BorderRadius.circular(9)),
//                             child: GestureDetector(
//                               onTap: () {
//                                 showDialog(
//                                     context: context,
//                                     builder: (context) => AlertDialog(
//                                           title: Text(
//                                             "Are you sure you want to logout ?",
//                                             style: GoogleFonts.poppins(
//                                               fontSize: 15,
//                                               color: Theme.of(context)
//                                                   .colorScheme
//                                                   .onSurface,
//                                             ),
//                                           ),
//                                           actions: [
//                                             TextButton(
//                                                 onPressed: () {
//                                                   Navigator.pop(context);
//                                                 },
//                                                 child: Text(
//                                                   "Cancel",
//                                                   style: TextStyle(
//                                                     color: Theme.of(context)
//                                                         .colorScheme
//                                                         .onSurface,
//                                                   ),
//                                                 )),
//                                             TextButton(
//                                                 onPressed: () {
//                                                   logOut();
//                                                   Navigator
//                                                       .pushNamedAndRemoveUntil(
//                                                           context,
//                                                           '/login',
//                                                           (route) => true);
//                                                 },
//                                                 child: Text("Log-Out"))
//                                           ],
//                                         ));
//                               },
//                               child: Container(
//                                 margin: EdgeInsets.only(top: 7),
//                                 child: Icon(
//                                   Icons.logout,
//                                   size: 30,
//                                   color:
//                                       Theme.of(context).colorScheme.onSurface,
//                                 ),
//                               ),
//                             )),
//                         SizedBox(
//                           width: MediaQuery.of(context).size.width * 0.05,
//                         ),
//                         Container(
//                             decoration: BoxDecoration(
//                                 borderRadius: BorderRadius.circular(9)),
//                             child: GestureDetector(
//                                 onTap: () {
//                                   Navigator.pushNamed(context, '/HomeChat');
//                                 },
//                                 child: Transform.rotate(
//                                   angle: 5.5,
//                                   child: Icon(
//                                     Icons.send_rounded,
//                                     size: 30,
//                                   ),
//                                 ))),
//                       ],
//                     ),
//                   ],
//                 ),





// --------logout button
                  // Container(
                  //     margin: EdgeInsets.only(right: 15),
                  //     child: GestureDetector(
                  //       onTap: () {
                  //         showDialog(
                  //             context: context,
                  //             builder: (context) => AlertDialog(
                  //                   title: Text(
                  //                     "Are you sure you want to logout ?",
                  //                     style: GoogleFonts.poppins(
                  //                       fontSize: 15,
                  //                       color: Theme.of(context)
                  //                           .colorScheme
                  //                           .onSurface,
                  //                     ),
                  //                   ),
                  //                   actions: [
                  //                     TextButton(
                  //                         onPressed: () {
                  //                           Navigator.pop(context);
                  //                         },
                  //                         child: Text(
                  //                           "Cancel",
                  //                           style: TextStyle(
                  //                             color: Theme.of(context)
                  //                                 .colorScheme
                  //                                 .onSurface,
                  //                           ),
                  //                         )),
                  //                     TextButton(
                  //                         onPressed: () {
                  //                           logOut();
                  //                           Navigator.pushNamedAndRemoveUntil(
                  //                               context,
                  //                               '/login',
                  //                               (route) => true);
                  //                         },
                  //                         child: Text("Log-Out"))
                  //                   ],
                  //                 ));
                  //       },
                  //       child: Container(
                  //         margin: EdgeInsets.only(top: 7),
                  //         child: Icon(
                  //           Icons.logout,
                  //           size: 30,
                  //           // color: Theme.of(context).colorScheme.onSurface,
                  //         ),
                  //       ),
                  //     )),