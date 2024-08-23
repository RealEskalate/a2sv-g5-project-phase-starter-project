import 'pages.dart';

class SearchPage extends StatelessWidget {
  const SearchPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        title: Row(
          children: [
            backButton(
              iconColor: const Color(
                0xFF3F51F3,
              ),
              onTap: () {
                Navigator.of(context).pop();
                context.read<ProductBloc>().add(LoadAllProductEvent());
              },
            ),
            const SizedBox(
              width: 70,
            ),
            const CustomText(
              text: 'Search Product',
              fontSize: 16,
            ),
          ],
        ),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            Row(
              children: [
                SizedBox(
                  width: MediaQuery.sizeOf(context).width * 0.75,
                  height: 48,
                  child: TextField(
                    onTapOutside: (event) {
                      FocusManager.instance.primaryFocus?.unfocus();
                    },
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(
                        borderSide: BorderSide(
                          color: Colors.grey,
                        ),
                        borderRadius: BorderRadius.all(
                          Radius.circular(
                            10,
                          ),
                        ),
                      ),
                      focusedBorder: OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.grey),
                        borderRadius: BorderRadius.all(
                          Radius.circular(
                            2,
                          ),
                        ),
                      ),
                      fillColor: Color.fromRGBO(253, 252, 252, 1),
                      filled: true,
                      suffixIcon: Icon(
                        Icons.arrow_forward,
                        color: Color(0xFF3F51F3),
                      ),
                      hintText: 'Search',
                    ),
                  ),
                ),
                const SizedBox(
                  width: 8,
                ),
                GestureDetector(
                  onTap: () {
                    showModalBottomSheet(
                      context: context,
                      builder: (BuildContext context) {
                        return Container(
                          height: 250,
                          padding: EdgeInsets.all(16),
                          child: Column(
                            children: [
                              const Align(
                                alignment: Alignment.topLeft,
                                child: CustomText(
                                  text: 'Category',
                                  fontSize: 16,
                                ),
                              ),
                              const SizedBox(
                                height: 10,
                              ),
                              Container(
                                height: 40,
                                decoration: BoxDecoration(
                                  border: Border.all(
                                    color: Colors.grey,
                                    width: 1,
                                  ),
                                  borderRadius: const BorderRadius.all(
                                    Radius.circular(8),
                                  ),
                                ),
                              ),
                              const SizedBox(
                                height: 10,
                              ),
                              const Align(
                                alignment: Alignment.topLeft,
                                child: CustomText(
                                  text: 'Price',
                                  fontSize: 16,
                                ),
                              ),
                              RangeSlider(
                                activeColor: const Color(0xFF3F51F3),
                                values: const RangeValues(
                                  0.2,
                                  0.8,
                                ),
                                onChanged: (value) {},
                              ),
                              const Spacer(),
                              const CustomOutlinedButton(
                                text: 'APPLY',
                                width: 300,
                                height: 44,
                                backgroundColor: Color(0xFF3F51F3),
                                color: Colors.white,
                              ),
                            ],
                          ),
                        );
                      },
                    );
                  },
                  child: Container(
                    width: 40,
                    height: 40,
                    decoration: const BoxDecoration(
                      color: Color(0xFF3F51F3),
                      borderRadius: BorderRadius.all(
                        Radius.circular(
                          2,
                        ),
                      ),
                    ),
                    child: const Icon(
                      Icons.filter_list,
                      color: Colors.white,
                    ),
                  ),
                )
              ],
            ),
            const SizedBox(
              height: 10,
            ),
            SizedBox(
              height: 590,
              child: ListView.builder(
                itemCount: 3,
                // TODO: implement the search
                itemBuilder: (context, idx) => const Placeholder(),
              ),
            )
          ],
        ),
      ),
    );
  }
}
