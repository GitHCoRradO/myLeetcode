class Solution {
	vector<int> twoSum(const vector<int> &input, int sum) {
		unordered_map<int, int> m;
		vector<int> result;
		for (int i = 0; i < input.size(); i++) {
			if (m.find(input[i]) == m.end()) {
      //insert numbers and indices into map
				m[sum - input[i]] = i;
        }
			else {
				result.push_back(m[input[i]] + 1);
				result.push_back(i + 1);
        //break is needed,for there is only one result in this case
        break;
			}
		return result;
}
