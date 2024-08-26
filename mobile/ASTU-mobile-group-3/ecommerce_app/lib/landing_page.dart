import 'package:flutter/material.dart';
import 'core/themes/themes.dart';
import 'features/chat/presentation/pages/chat_list.dart';
import 'features/product/presentation/pages/add_product_page.dart';
import 'features/product/presentation/pages/product_list_page.dart';

class LandingPage extends StatefulWidget {
  static const String routes = '/landing_page';
  // final bool isLogedIn;
  const LandingPage({
    super.key,
  });

  @override
  State<LandingPage> createState() => _LandingPageState();
}

class _LandingPageState extends State<LandingPage> {
  late bool isLogedIn;
  int currentTab = 0;

  final PageStorageBucket bucket = PageStorageBucket();
  Widget currentScreen = ProductListPage();

  @override
  Widget build(BuildContext context) {
    final List<Widget> screens = [
      ProductListPage(),
      const ChatList(),
    ];
    return Scaffold(
      resizeToAvoidBottomInset: false,
      backgroundColor: Theme.of(context).colorScheme.tertiary,
      body: PageStorage(
        bucket: bucket,
        child: screens[currentTab],
      ),
      floatingActionButton: FloatingActionButton(
        backgroundColor: MyTheme.ecBlue,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(100),
        ),
        onPressed: () {
          Navigator.pushNamed(context, AddProductPage.routes);
        },
        child: const Icon(
          Icons.add,
          color: MyTheme.ecWhite,
        ),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.centerDocked,
      bottomNavigationBar: BottomAppBar(
        elevation: 10,
        color: const Color.fromARGB(255, 13, 32, 58),
        shape: const CircularNotchedRectangle(),
        child: SizedBox(
          height: 50.0,
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: <Widget>[
              GestureDetector(
                onTap: () {
                  setState(() {
                    currentTab = 0;
                  });
                },
                child: currentTab == 0
                    ? Container(
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(10),
                          color: const Color.fromARGB(255, 73, 73, 80),
                        ),
                        child: const Padding(
                          padding: EdgeInsets.all(10),
                          child: Icon(
                            Icons.list,
                            color: MyTheme.ecWhite,
                          ),
                        ),
                      )
                    : Container(
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(10),
                          color: Colors.transparent,
                        ),
                        child: const Padding(
                          padding: EdgeInsets.all(20.0),
                          child: Icon(
                            Icons.list,
                            color: MyTheme.ecWhite,
                          ),
                        ),
                      ),
              ),
              GestureDetector(
                onTap: () {
                  setState(() {
                    currentTab = 1;
                  });
                },
                child: currentTab == 1
                    ? Container(
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(10),
                          color: MyTheme.ecBlue,
                        ),
                        child: const Padding(
                          padding: EdgeInsets.all(10),
                          child: Icon(
                            Icons.chat,
                            color: MyTheme.ecWhite,
                          ),
                        ),
                      )
                    : Container(
                        decoration: BoxDecoration(
                            borderRadius: BorderRadius.circular(10),
                            color: Colors.transparent),
                        child: const Padding(
                          padding: EdgeInsets.all(20.0),
                          child: Icon(
                            Icons.chat,
                            color: MyTheme.ecWhite,
                          ),
                        ),
                      ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
